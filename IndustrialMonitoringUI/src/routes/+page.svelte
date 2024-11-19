<script lang="ts">
    import {goto} from '$app/navigation';

    function goToOrg() {
        goto('/private/org/1');
    }
    export let data;
	$: ({ countries } = data)
</script>

<h1>Welcome to your library project</h1>
<p>Create your package using @sveltejs/package and preview/showcase your work with SvelteKit</p>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<button on:click={goToOrg}>Go to Org 1</button>

<h1 class="text-3xl font-bold underline">Welcome to Supabase!</h1>
<ul>
	{#each countries as country}
		<li>{country.name}</li>
	{/each}
</ul>