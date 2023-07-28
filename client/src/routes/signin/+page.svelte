<script lang="ts">
	let email: string;
	let password: string;
</script>

<div class="flex h-screen">
	<div class="m-auto drop-shadow-xl bg-white rounded py-2 px-5 text-slate-700">
		<h1 class="text-2xl font-bold bg-slate-200 w-fit px-3 py-1 rounded">Sign in</h1>
		<div class="my-5">
			<div class="my-2 w-72">
				<label for="email">
                    Email
                    <input bind:value={email} class="border-2 rounded-sm w-full px-1 my-2 h-9" type="text" />
                </label>
                <label for="password">
                    Password
                    <input bind:value={password} class="border-2 rounded-sm w-full px-1 my-2 h-9" type="text"/>
                </label>
			</div>
			<button class="font-bold w-full btn btn-success rounded">SIGN IN</button>
		</div>
		<div class="text-center">
			<a class="text-slate-700/50" href="/signup">don't have an account ?</a>
		</div>
	</div>
</div>
