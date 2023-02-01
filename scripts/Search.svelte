<script lang="ts">
    type SearchResult = {
        id: number
        firstName: string
        lastName: string
        slug: string
        rating: number
    }

    // Composers found
    let composers: SearchResult[] = []
    // Tuples with search queries: the currently being requested from API, and the last inputted
    let queryStack: [string | undefined, string | undefined] = [undefined, undefined]
    // Should search be shown
    let searchVisible = false
    // Selected search result
    let selectedResult = 0

    /**
     * Performs request to API
     * @param query Search query
     */
    async function getFromApi(query: string | undefined): Promise<SearchResult[]> {
        try {
            const response = await fetch(`/api/search?q=${query}`)
            if (!response.ok) {
                return []
            }
            return await response.json()
        } catch (err: unknown) {
            console.log(err)
            return []
        }
    }

    /** Sends queries to API. Implements throttling so that no more than one parallel search request can be executed at each given moment. */
    async function queryApi() {
        if (queryStack[0] !== undefined) {
            composers = await getFromApi(queryStack[0])
            if (queryStack[1] !== undefined) {
                // Send the latest query to the API after the current one is done
                queryStack[0] = queryStack[1]
                queryStack[1] = undefined
                await queryApi()
            } else {
                queryStack[0] = undefined
            }
        } else {
            // Clean up results when there are no more queries
            composers = []
        }
    }

    /** Shows search input and results */
    function showSearch(): void {
        searchVisible = true
    }

    /** Hides search input and results */
    function hideSearch(): void {
        composers = []
        searchVisible = false
    }

    /** Focuses on given element */
    function focus(el: HTMLElement): void {
        el.focus()
    }

    /** Dispatches actions based on keys pressed inside search input */
    function handleKeydown(event: KeyboardEvent) {
        if (event.code === 'ArrowUp' && composers.length > 0) {
            selectedResult = selectedResult > 0
                ? selectedResult - 1
                : composers.length - 1
        } else if (event.code === 'ArrowDown') {
            selectedResult = selectedResult < composers.length - 1
                ? selectedResult + 1
                : selectedResult = 0
        } else if (event.code === 'Escape') {
            hideSearch()
        } else if (event.code === 'Enter' && composers.length > 0) {
            location.pathname = `/composer/${composers[selectedResult].slug}`
        }
    }

    /** Executes given handler when user clicked outside given element */
    function clickOutside(
        node: HTMLElement,
        handler: () => void
    ): { destroy: () => void } {
        const onClick = (event: MouseEvent) =>
            node &&
            !node.contains(event.target as HTMLElement) &&
            !event.defaultPrevented &&
            handler();

        document.addEventListener('click', onClick, true);

        return {
            destroy() {
                document.removeEventListener('click', onClick, true);
            },
        };
    }

    /**
     * Saves user input into search field and sends requests to API
     * @param event User input
     */
    function handleSearch(event: { target: HTMLInputElement; }): void {
        // Convert empty string to undefined so it's more clear
        const inputEvent = event.target.value || undefined
        if (queryStack[0] === undefined) {
            // If no current searches, start the search
            queryStack[0] = inputEvent
            queryApi()
        } else {
            // If there is active search, save current input for the next search to run
            queryStack[1] = inputEvent
        }
    }

    function handleResultHover(ind: number): void {
        selectedResult = ind
    }
</script>


<div class="search-button" on:click={showSearch} on:keypress={showSearch}>
    <img src="/static/img/search-icon.svg" alt="Search"/>
</div>

{#if searchVisible}
    <div class="search-wrapper" on:keydown={handleKeydown}>
        <div class="search"
             use:clickOutside={hideSearch}>
            <input class="search__field"
                   type="search"
                   placeholder="Search composers by last name"
                   on:input={handleSearch}
                   use:focus/>
            {#each composers as composer, i}
                <a href="/composer/{composer.slug}" on:mouseenter={() => handleResultHover(i)}>
                    <div class:search__result_selected="{selectedResult === i}" class="search__result">
                        {composer.lastName}, {composer.firstName}
                    </div>
                </a>
            {/each}
        </div>
    </div>
{/if}