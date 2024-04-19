<script>
  import EndpointStatusBox from "../lib/EndpointStatusBox.svelte";

  export let endpointName = "endpoint1";
  export let URL = "/api/example";
  export let status = "Up";
  export let loginState = "Login";
</script>

<link
  rel="stylesheet"
  href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
/>

<div class="status-evolution">
  <div class="card border-secondary" style="width: 20rem;">
    <div class="card-body">
      <h2 class="card-title" style="margin-bottom: 12px">{endpointName}</h2>
      <hr />
      <p class="card-text">{URL}</p>
      <p class="card-text" style="margin-bottom: 8px">History:</p>
      <div class="endpoint-container">
        <EndpointStatusBox color="green" />
        <EndpointStatusBox color="green" />
        <EndpointStatusBox color="green" />
        <EndpointStatusBox color="green" />
        <EndpointStatusBox />
        <EndpointStatusBox />
        <EndpointStatusBox color="red" />
        <EndpointStatusBox />
        <EndpointStatusBox />
        <EndpointStatusBox color="green" />
      </div>
      <p class="card-text">Status: {status}</p>
      <p class="card-text" style="margin-bottom: 0">Bugs:</p>
      <ul>
        <li>Bug message</li>
        <li>Bug message 2</li>
        <li>Bug message 3</li>
        <li>Bug message 4</li>
      </ul>
      {#if loginState === "Login"}
        <div class="final-button">
          <div class="btn-group dropend">
            <button
              type="button"
              class="btn btn-danger dropdown-toggle"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              style="height: 40px; text-align: center; display: flex; align-items: center; justify-content: center;"
            >
              Report a bug
            </button>

            <form class="dropdown-menu p-3" style="width: 300px">
              <div class="mb-3">
                <label for="appName" class="form-label">Message:</label>
                <input
                  type="text"
                  class="form-control"
                  id="appName"
                  placeholder="Example app"
                />
              </div>
              <button type="submit" class="btn btn-primary">Submit bug</button>
            </form>
          </div>
        </div>
      {:else}
        <div class="final-button">
          <button
            type="button"
            class="btn btn-danger"
            style="display: flex; align-items: center; justify-content: center; height: 40px; width: 40px ;"
            ><span
              class="material-symbols-outlined"
              style="height: fit-content"
            >
              delete
            </span></button
          >
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .status-evolution {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(20rem, 1fr));
  }
  .card {
    margin-bottom: 32px;
  }
  .endpoint-container {
    display: flex;
    flex-direction: row;
  }

  .material-symbols-outlined {
    font-variation-settings:
      "FILL" 0,
      "wght" 400,
      "GRAD" 0,
      "opsz" 24;
  }
  .final-button {
    display: flex;
    justify-content: end;
  }
</style>
