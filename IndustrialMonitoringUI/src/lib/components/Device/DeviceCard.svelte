<script lang="ts">
   import { createEventDispatcher } from 'svelte';
    import type { Device } from '$lib/models/Device.js';
    import {goto} from '$app/navigation';
    import { page } from '$app/stores';

  export let device: Device;
  
  const handleDeviceClick = (deviceId: number) => {
    const orgId = $page.params.orgId;
    console.log(orgId)
    goto(`private/device_manager/${deviceId}`);
  };

  </script>

<main>
  <div class="card" on:click={() => {handleDeviceClick(device.id)}}>
    <h3>{device.name}</h3>
    <p>{device.bucket}</p>
</div>
</main>

<style>
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
      background-color: #e0e0e0;
    }
</style>