<script lang="ts">
    import DeviceList from "$lib/components/Device/DeviceList.svelte";
    import { page } from '$app/stores';
    import { get } from 'svelte/store';
    const { orgId } = get(page).params;
</script>

<main>
  <h1>Device Information</h1>
  <DeviceList id={orgId} />
</main>