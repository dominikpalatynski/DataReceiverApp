<script lang="ts">
  import { onMount } from 'svelte';
  import { writable } from 'svelte/store';
  import AddDevice from './AddDevice.svelte';
    import DeviceCard from './DeviceCard.svelte';
    import type { Device } from '$lib/models/Device.ts';
    import devicesStore from '$lib/stores/devicesStore.js';
    import { fetchDevicesByOrgId, addDeviceInfo } from '$lib/services/deviceService.js';
    import {goto} from '$app/navigation';
    import { page } from '$app/stores';

  export let id: string;

  let loading = writable(true);
  let error = writable(null);
  let showForm = false;

  const fetchDevices = async () => {
    try {
        const data: Device[] = await fetchDevicesByOrgId(id)
        devicesStore.set(data)
    }
    catch (err) {
        console.error(err)
        loading.set(false)
    }
    finally {
        loading.set(false)
    }
  };

  const addDevice = async (device: { name: string; interval: number; org_id: number}) => {
    try {
        await addDeviceInfo(device);
        await fetchDevices();
    } catch (err) {
      console.error(err);
      error.set(err.message);
    }
  };

  const toggleFormVisibility = () => {
    showForm = !showForm;
  };

  onMount(fetchDevices);

</script>

<main>
<button on:click={toggleFormVisibility} class="add-button">
    {showForm ? 'Ukryj formularz' : 'Dodaj urządzenie (+)'}
    </button>

    {#if showForm}
    <AddDevice onAddDevice={addDevice} />
    {/if}

  {#if $loading}
    <p>Loading...</p>
  {:else if $error}
    <p>Error: {$error}</p>
    {:else if $devicesStore.length > 0}
    <div class="grid">
      {#each $devicesStore as device}
        <DeviceCard {device}/>
      {/each}
    </div>
  {:else}
    <p>No devices found.</p>
  {/if}
</main>

<style>
    main {
      font-family: Arial, sans-serif;
      text-align: center;
      padding: 20px;
    }
  
    .grid {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
      gap: 20px;
    }
  
    .card {
      background-color: #f9f9f9;
      border: 1px solid #ddd;
      border-radius: 8px;
      padding: 20px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
      transition: transform 0.3s, box-shadow 0.3s;
      cursor: pointer;
    }
  
    .card:hover {
      transform: translateY(-4px);
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
      background-color: #e0e0e0; /* Podświetlenie */
    }
    .add-button {
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    padding: 10px 20px;
    cursor: pointer;
    margin-bottom: 20px;
  }

  .add-button:hover {
    background-color: #0056b3;
  }

</style>