<script lang="ts">
    import { addSensor, fetchSlotsByDeviceId } from '$lib/services/deviceService.js';
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    // Typowanie danych
    interface Slot {
      device_id: number;
      slot_number: number;
      sensor_id?: number | null;
    }
  
    let slots: Slot[] = [];
    let selectedSlot: Slot | null = null;
    let sensorId: string = '';
    let sensorName: string = '';
    let variableName: string = '';

    // Pobieranie listy slotów z backendu
    const fetchSlots = async (): Promise<void> => {
        const deviceId = $page.params.deviceId;
        try {
            slots = await fetchSlotsByDeviceId(deviceId);
            slots = slots.sort((a, b) => a.slot_number - b.slot_number);
        } catch (error) {
            console.error('Error:', error);
        }
    };
  
    onMount(fetchSlots);
  
    const handleEditSlot = (slot: Slot): void => {
      selectedSlot = slot;
      sensorId = slot.sensor_id?.toString() ?? '';
    };
  
    // Obsługa aktualizacji slotu
    const handleUpdateSlot = async (): Promise<void> => {
      if (!selectedSlot) return;
  
      try {
            const sensor = await addSensor({
                name: sensorName,
                variable_name: variableName,
                device_id: selectedSlot.device_id,
                slot: selectedSlot.slot_number
            });
            selectedSlot = null;
            await fetchSlots();
          
      } catch (error) {
        console.error('Error:', error);
      }
    };
  </script>
  
  <h1 class="text-2xl font-bold text-center mt-8">Slot List</h1>
  
  <div class="slot-list flex flex-col gap-4 max-w-md mx-auto mt-6">
    {#each slots as slot}
      <div
        class="slot-item p-4 border border-gray-300 rounded cursor-pointer hover:bg-gray-100"
        on:click={() => handleEditSlot(slot)}
      >
        <p>Slot Number: {slot.slot_number}</p>
        <p>Sensor ID: {slot.sensor_id ?? 'N/A'}</p>
        {#if !slot.sensor_id}
        <button
          on:click={(e) => { e.stopPropagation(); handleAddSensor(slot); }}
          class="mt-2 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition"
        >
          Add Sensor
        </button>
      {/if}
      </div>
    {/each}
    {#if selectedSlot}
    <div class="edit-form max-w-md mx-auto mt-8 p-6 bg-gray-50 rounded shadow">
      <h2 class="text-xl font-semibold mb-4">Edit Slot</h2>
      <label class="block mb-2">
        <span class="text-gray-700">Sensor Name:</span>
        <input
          type="text"
          bind:value={sensorName}
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring focus:ring-blue-200"
          placeholder="Enter sensor name"
        />
      </label>
  
      <!-- Pole do wpisania nazwy zmiennej -->
      <label class="block mb-2">
        <span class="text-gray-700">Variable Name:</span>
        <input
          type="text"
          bind:value={variableName}
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring focus:ring-blue-200"
          placeholder="Enter variable name"
        />
      </label>
      <div class="flex justify-end gap-4 mt-4">
        <button
          on:click={handleUpdateSlot}
          class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
        >
          Save
        </button>
        <button
          on:click={() => (selectedSlot = null)}
          class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400 transition"
        >
          Cancel
        </button>
      </div>
    </div>
  {/if}

  </div>
  
 
  