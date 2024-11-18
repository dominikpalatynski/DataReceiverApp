<script lang="ts">
  	import { Input } from "$lib/components/ui/input/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { createOrg } from "$lib/services/orgService.js";

    const onCreateOrg = async (org: { name: string; bucket: string;}) => {
    try {
        await createOrg(org);
    } catch (err) {
        console.error(err);
        error.set(err.message);
    }
    };

    let name = '';
    let bucket = '';
  
    const handleSubmit = (e: Event) => {
      e.preventDefault();
      console.log(name, bucket);
      if (name && bucket) {
        onCreateOrg({ name, bucket});
        name = '';
        bucket = '';
      }
    };
  </script>
  
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
  <form  on:submit={handleSubmit} class = "bg-white rounded-lg shadow-lg p-8 max-w-md w-full space-y-6">
    <h2 class="text-2xl font-bold text-gray-800 text-center">Stwórz organizację</h2>

    <div class="space-y-2">
      <label for="name" class="text-black-600 font-medium">Nazwa organizacji</label>
      <Input
        id="name"
        name="name"
        type="name"
        placeholder="nazwa organizacji"
        class="w-full text-gray-800"
        bind:value={name}
        required
      />
    </div>

    <div class="space-y-2">
        <label for="bucket" class="text-black-600 font-medium">Nazwa bucketu</label>
        <Input
          id="bucket"
          name="bucket"
          type="bucket"
          placeholder="nazwa bucketu"
          class="w-full text-gray-800"
          bind:value={bucket}
          required
        />
      </div>

    <div class="space-y-4">
      <Button type="submit" class="w-full bg-blue-600 text-white font-semibold hover:bg-blue-700 transition duration-200">
        Create
      </Button>
    </div>
  </form>
</div>