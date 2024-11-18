<script lang="ts">
    import { page } from '$app/stores';
    import { get } from 'svelte/store';

  export let onAddDevice: (device: { name: string; interval: number; org_id: number }) => void;

  let name = '';
  let interval = 1; // domyślny interwał
  const handleSubmit = (e: Event) => {
    e.preventDefault();
    if (name && interval > 0) {
      const { orgId } = get(page).params;

      onAddDevice({name, interval, org_id: parseInt(orgId)});
      name = '';
      interval = 1;
    }
  };
</script>

<form on:submit={handleSubmit}>
  <label>
    Nazwa urządzenia:
    <input type="text" bind:value={name} required />
  </label>
  <label>
    Interwał:
    <input type="number" bind:value={interval} min="1" required />
  </label>
  <button type="submit">Dodaj urządzenie</button>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;
    max-width: 400px;
    margin: auto;
    gap: 10px;
    padding-bottom: 20px;
  }

  label {
    display: flex;
    flex-direction: column;
  }

  button {
    padding: 10px;
    background-color: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  button:hover {
    background-color: #218838;
  }
</style>