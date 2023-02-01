import Search from './Search.svelte'

const target = document.getElementById('searchBlock')
const app = target ? new Search({target}) : undefined

export default app