<script lang="ts">
  import { onMount } from 'svelte';
  import {goto} from '$app/navigation';
  import { fetchOrgsConnectedToUser } from "$lib/services/orgService.js";
    import type { UserOrganization } from '$lib/models/Org.js';

  let organizations: UserOrganization[] = [];
  let error = '';

  const fetchDevices = async () => {
    try {
        organizations = await fetchOrgsConnectedToUser()
    }
    catch (err) {
        console.error(err)
    }
  };

  const handleDeviceClick = (orgId: number) => {
    goto(`org/${orgId}`);
  };

  onMount(fetchDevices);


</script>

<h2 class="text-2xl font-bold text-gray-800 mb-6">Twoje Organizacje</h2>

{#if error}
  <p class="text-red-600 font-semibold">{error}</p>
{:else if organizations.length === 0}
  <p class="text-gray-500">Nie znaleziono organizacji powiązanych z użytkownikiem.</p>
{:else}
  <div class="space-y-4">
    {#each organizations as org}
      <div on:click={() => {handleDeviceClick(org.organization.id)}} class="bg-white shadow-lg rounded-lg p-6 border border-gray-200">
        <h3 class="text-xl font-semibold text-gray-700">{org.organization.name}</h3>
        <p class="text-gray-600 mt-2">
          ID organizacji: <span class="text-gray-800 font-medium">{org.organization.id}</span>
        </p>
        <p class="text-gray-600 mt-1">
          Rola: <span class="text-indigo-600 font-semibold">{org.role}</span>
        </p>
      </div>
    {/each}
  </div>
{/if}