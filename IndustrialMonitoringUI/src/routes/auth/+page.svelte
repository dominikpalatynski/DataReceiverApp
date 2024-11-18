<script lang="ts">
	import { Input } from "$lib/components/ui/input/index.js";
	import { Button } from "$lib/components/ui/button/index.js";
  </script>
   
 

<!-- <form method="POST" action="?/login">
	<label>
		Email
		<Input name="email" type="email"/>
	</label>
	<label>
		Password
		<Input  name="password" type="password"/>
	</label>
	<button>Login</button>
	<button formaction="?/signup">Sign up</button>
</form> -->

<div class="flex items-center justify-center min-h-screen bg-gray-100">
	<form method="POST" action="?/login" class="bg-white rounded-lg shadow-lg p-8 max-w-md w-full space-y-6">
	  <h2 class="text-2xl font-bold text-gray-800 text-center">Sign in to your account</h2>
  
	  <div class="space-y-2">
		<label for="email" class="text-gray-600 font-medium">Email</label>
		<Input
		  id="email"
		  name="email"
		  type="email"
		  placeholder="you@example.com"
		  class="w-full"
		  required
		/>
	  </div>
  
	  <div class="space-y-2">
		<label for="password" class="text-gray-600 font-medium">Password</label>
		<Input
		  id="password"
		  name="password"
		  type="password"
		  placeholder="••••••••"
		  class="w-full"
		  required
		/>
	  </div>
  
	  <div class="space-y-4">
		<Button type="submit" class="w-full bg-blue-600 text-white font-semibold hover:bg-blue-700 transition duration-200">
		  Login
		</Button>
		<Button
		  type="submit"
		  formaction="?/signup"
		  class="w-full bg-gray-600 text-white font-semibold hover:bg-gray-700 transition duration-200"
		>
		  Sign up
		</Button>
	  </div>
  
	  <div class="flex justify-between text-sm text-gray-500 pt-4">
		<Button class="hover:underline">Forgot your password?</Button>
		<Button class="hover:underline">Create an account</Button>
	  </div>
	</form>
  </div>