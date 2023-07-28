<script lang="ts">
	type Gender = 'male' | 'female' | 'other';

	let username: string;
	let email: string;
	let password: string;
	let confirmPassword: string;
	let age: string;
	let gender: Gender;
</script>

<div class="flex h-screen">
	<div class="m-auto drop-shadow-xl bg-white rounded py-2 px-5 text-slate-700">
		<h1 class="text-2xl font-bold bg-slate-200 w-fit px-3 py-1 rounded">Sign up</h1>
		<div class="my-5">
			<div class="my-2 w-72">
				<label for="username">
					Username
					<input
						bind:value={username}
						class="border-2 rounded-sm w-full px-1 my-2 h-9"
						type="text"
					/>
				</label>
				<label for="email">
					Email
					<input bind:value={email} class="border-2 rounded-sm w-full px-1 my-2 h-9" type="text" />
				</label>
				<label for="password">
					Password
					<input
						bind:value={password}
						class="border-2 rounded-sm w-full px-1 my-2 h-9"
						type="text"
					/>
				</label>
				<label for="confirm-password">
					Confirm Password
					<input
						bind:value={confirmPassword}
						class="border-2 rounded-sm w-full px-1 my-2 h-9"
						type="text"
					/>
				</label>
				<label for="age">
					Age
					<input bind:value={age} class="border-2 rounded-sm w-full my-2 px-1 h-9" type="text" />
				</label>
				<div class="flex gap-2 my-2 justify-between">
					<button
						on:click={() => (gender = 'male')}
						class="rounded w-full border-2 p-1 flex justify-around"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="male"
							checked={gender == 'male'}
						/>
						<span class="w-full">Male</span>
					</button>
					<button
						on:click={() => (gender = 'female')}
						class="rounded w-full border-2 p-1 flex justify-around"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="female"
							checked={gender == 'female'}
						/>
						<span class="w-full">Female</span>
					</button>
					<button
						on:click={() => (gender = 'other')}
						class="rounded w-full border-2 p-1 flex justify-around"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="other"
							checked={gender == 'other'}
						/>
						<span class="w-full">Other</span>
					</button>
				</div>
			</div>
			<button class="font-bold w-full btn btn-success rounded">SIGN UP</button>
		</div>
		<div class="text-center text-slate-400">
			<a class="text-slate-700/50" href="/signin">already have an account ?</a>
		</div>
	</div>
</div>
