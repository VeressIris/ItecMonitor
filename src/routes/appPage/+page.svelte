<script>
  import EndpointCard from "../../lib/EndpointCard.svelte";
  let appName = "Example app";
  let endpoints = [
    {
      name: "Endpoint 1",
      url: "/api/endpoint1",
      status: "Up",
    },
    {
      name: "Endpoint 2",
      url: "/api/endpoint2",
      status: "Down",
    },
    {
      name: "Endpoint 3",
      url: "/api/endpoint3",
      status: "Up",
    },
  ];

  function getEndpoints() {}

  getEndpoints();
</script>

<div class="header">
  <h1>{appName}</h1>
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
