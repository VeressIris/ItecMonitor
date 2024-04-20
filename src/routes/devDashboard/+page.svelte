<script>
  import AppCard from "../../lib/AppCard.svelte";

  let appName = "";
  let baseURL = "";

  let apps = [
    {
      name: "Example app",
      baseURL: "https://example.com/",
      endpoints: 3,
      status: "down",
    },
    {
      name: "Another app",
      baseURL: "https://another.com/",
      endpoints: 5,
      status: "up",
    },
  ];

  function addApp() {
    const newApp = {
      name: appName,
      baseURL: baseURL,
      endpoints: 0,
      status: "Down",
    };
    apps = [...apps, newApp];
    // addAppToDB(newApp);
  }

  function addAppToDB(app) {
    fetch("http://127.0.0.1:3000/addApp", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(app),
    })
      .then((response) => {
        console.log(response);
        response.json();
      })
      .catch((error) => {
        console.error("There was a problem with the fetch operation:", error);
      });
  }
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
          bind:value={appName}
        />
      </div>
      <div class="mb-3">
        <label for="baseurl" class="form-label">BaseURL:</label>
        <input
          type="text"
          class="form-control"
          id="baseurl"
          placeholder="https://example.com/"
          bind:value={baseURL}
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
