<script>
  import { onMount } from "svelte";
  let test = "";
  onMount(async () => {
    const response = await fetch("http://127.0.0.1:3000");
    test = await response.text();
  });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
  Go says: {test}
</p>
