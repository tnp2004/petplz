<script lang="ts">
	import type { PageData } from "./$types";

    export let data: PageData;
    const id = data.params.slug
</script>

<div class="w-2/3 mx-auto">
    <div class="card lg:card-side bg-base-100 shadow-xl">
        <figure class="w-1/3 bg-slate-100"><img src="https://daisyui.com/images/stock/photo-1494232410401-ad00d5433cfa.jpg" alt="Album"/></figure>
        <div class="card-body">
          <h2 class="card-title">{id}</h2>
          <p>Click the button to listen on Spotiwhy app.</p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary">Listen</button>
          </div>
        </div>
    </div>
</div>