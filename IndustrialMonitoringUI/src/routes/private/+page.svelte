<script lang="ts">
	import { invalidate } from '$app/navigation';
	import type { EventHandler } from 'svelte/elements';

	import type { PageData } from './$types.js'

	export let data: PageData;
	$: ({user } = data);

</script>

<h1>Private page for user: {user?.email}</h1>
<h2>Notes</h2>