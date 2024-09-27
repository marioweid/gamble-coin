<script lang="ts">
	interface Beer {
		name: string;
		price: string;
	}
    let disabled = false;
	let promise: Promise<Beer[]> = Promise.resolve([]);

	async function fetchBeers(): Promise<Beer[]> {
		const response = await fetch('/api/beers', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		if (response.ok) {
            return response.json();
		} else {
            throw new Error('Failed to fetch data from backend');
        }
	}

    function handleClick(){
        promise = fetchBeers();
        disabled = true;
    }
</script>

<button
	class="flex items-center px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
	on:click={() => handleClick()}
    { disabled }
>
	<span class="mx-1">Fetch Beers from Go-Backend</span>
</button>

{#await promise}
	<p>Loading</p>
{:then beers}
    {#if beers.length > 0}
        <p><b>Beer</b> - <b>Price</b></p>    
    {/if}
	{#each beers as beer}
		<p>{beer.name} - {beer.price}</p>
	{/each}
{:catch error}
	<p>Error happened! {error}</p>
{/await}
