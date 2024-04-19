<script>
  import AppCard from "../../lib/AppCard.svelte";

  let apps = [
    {
      name: "Example app",
      baseURL: "https://example.com/",
      endpoints: 3,
      Status: "down",
    },
    {
      name: "Another app",
      baseURL: "https://another.com/",
      endpoints: 5,
      Status: "up",
    },
  ];

  function addApp(event) {}
</script>

<div class="header">
  <h1>Developer Dashboard</h1>
  <div class="btn-group dropend">
    <button
      type="button"
      class="btn btn-success dropdown-toggle"
      data-bs-toggle="dropdown"
      aria-expanded="false"
      style="height: 40px; text-align: center; display: flex; align-items: center; justify-content: center;"
    >
      Add an app
    </button>
    <form
      class="dropdown-menu p-3"
      style="width: 300px"
      on:submit|preventDefault={addApp}
    >
      <div class="mb-3">
        <label for="appName" class="form-label">App name:</label>
        <input
          type="text"
          class="form-control"
          id="appName"
          placeholder="Example app"
        />
      </div>
      <div class="mb-3">
        <label for="baseurl" class="form-label">BaseURL:</label>
        <input
          type="text"
          class="form-control"
          id="baseurl"
          placeholder="https://example.com/"
        />
      </div>
      <button type="submit" class="btn btn-primary">Add app</button>
    </form>
  </div>
</div>

<div class="list">
  {#each apps as app}
    <AppCard
      appName={app.name}
      baseURL={app.baseURL}
      endpoints={app.endpoints}
      status={app.status}
    />
  {/each}
</div>

<style>
  .header h1 {
    margin-right: 32px;
  }
</style>
