<script lang="ts">
	let email: string;
	let password: string;
</script>

<div class="flex h-screen">
	<div class="m-auto drop-shadow-xl bg-white rounded py-2 px-5 text-slate-700">
		<h1 class="text-2xl font-bold">Sign in</h1>
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
			<button class="bg-emerald-200 font-bold w-full rounded-sm p-1 h-9">SIGN IN</button>
		</div>
		<div class="text-center text-slate-400">
			<a href="/signup">Don't have an account yet ?</a>
		</div>
	</div>
</div>
