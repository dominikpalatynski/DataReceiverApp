<script>
		export let data;
	$: ({ supabase } = data);
	
	let organizations = ["Org 1", "Org 2", "Org 3"];
	let showOrgDropdown = false; // Flaga dla menu rozwijanego organizacji
	let showSettingsDropdown = false; // Flaga dla menu rozwijanego ustawień
	
	function toggleOrgDropdown() {
	  showOrgDropdown = !showOrgDropdown;
	}
	
	function toggleSettingsDropdown() {
	  showSettingsDropdown = !showSettingsDropdown;
	}
  </script>
  
  <style>
	.sidebar {
	  width: 150px;
	}
  </style>

  <div class="fixed top-0 left-0 h-screen bg-gray-800 text-white sidebar flex flex-col p-4 space-y-6 shadow-lg">
	<!-- Logo -->
	<div class="text-2xl font-bold mb-6">App Name</div>
  
	<!-- Organizacje - Dropdown -->
	<div>
	  <button
		on:click={toggleOrgDropdown}
		class="w-full flex justify-between items-center bg-gray-700 px-4 py-2 rounded-md hover:bg-gray-600"
	  >
		<span>Wybierz organizację</span>
		<svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
		  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
		</svg>
	  </button>
	  {#if showOrgDropdown}
		<div class="mt-2 bg-gray-700 rounded-md shadow-md">
		  <ul class="text-gray-300">
			{#each organizations as org}
			  <li class="px-4 py-2 hover:bg-gray-600 cursor-pointer">{org}</li>
			{/each}
			<li class="px-4 py-2 hover:bg-gray-600 cursor-pointer text-green-400">
			  + Dodaj nową organizację
			</li>
		  </ul>
		</div>
	  {/if}
	</div>
  
	<!-- Ustawienia - Dropdown -->
	<div>
	  <button
		on:click={toggleSettingsDropdown}
		class="w-full flex justify-between items-center bg-gray-700 px-4 py-2 rounded-md hover:bg-gray-600"
	  >
		<span>Ustawienia</span>
		<svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
		  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
		</svg>
	  </button>
	  {#if showSettingsDropdown}
		<div class="mt-2 bg-gray-700 rounded-md shadow-md">
		  <ul class="text-gray-300">
			<li class="px-4 py-2 hover:bg-gray-600 cursor-pointer">Profil</li>
			<li class="px-4 py-2 hover:bg-gray-600 cursor-pointer">Bezpieczeństwo</li>
			<li class="px-4 py-2 hover:bg-gray-600 cursor-pointer">Wyloguj się</li>
		  </ul>
		</div>
	  {/if}
	</div>
  </div>
  
  <main class="ml-[250px] p-6">
	<slot />
  </main>
