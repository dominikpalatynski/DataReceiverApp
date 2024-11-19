<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { get } from 'svelte/store';
    import type { Device } from '$lib/models/Device.js';
    import { fetchDeviceById } from '$lib/services/deviceService.js';
  
    let device: Device | null = null;
    let loading = true;
    let error = null;
  
    const { deviceId } = get(page).params;
  
    onMount(async () => {
      try {
        device = await fetchDeviceById(deviceId);
        loading = false;
      } catch (err) {
        error = err.message;
        loading = false;
      }
    });
  </script>
  
  <main>
    {#if loading}
      <p>Loading device details...</p>
    {:else if error}
      <p>Error: {error}</p>
    {:else if device}
      <div class="device-details">
        <h2>Device: {device.name}</h2>
        <p><strong>Bucket:</strong> {device.bucket}</p>
        <p><strong>Interval:</strong> {device.interval} seconds</p>
        <p><strong>Organization ID:</strong> {device.org_id}</p>
      </div>
    {/if}
  </main>

  <!-- <main>
    <h1>Device Manager</h1>
  </main> -->