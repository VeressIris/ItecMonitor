<script>
  import { onMount } from "svelte";
  import EndpointCard from "../../../lib/EndpointCard.svelte";
  export let data;
  let endpoints = [];
  function getEndpoints() {
    console.log("Getting endpoints for: ", data.appName);
    fetch("http://127.0.0.1:3000/getEndpoints", {
      method: "POST",
      headers: {
        "Content-Type": "text/plain",
      },
      body: data.appName,
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        endpoints = data;
      })
      .catch((error) => {
        console.error(
          "There was a problem with the fetching the endpoints:",
          error
        );
      });
  }

  onMount(() => {
    getEndpoints();
  });
</script>

<div class="header">
  <h1>{data.appName}</h1>
</div>

<div class="list">
  {#each endpoints as endpoint}
    <EndpointCard
      endpointName={endpoint.name}
      URL={endpoint.url}
      status={endpoint.status}
    />
  {/each}
</div>
