<script lang="ts">
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "/vite.svg";
  import goLogo from "./assets/golang.svg";
  import Counter from "./lib/Counter.svelte";
</script>

<main>
  <div>
    <a href="https://vitejs.dev" target="_blank" rel="noreferrer">
      <img src={viteLogo} class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
    <a href="https://go.dev/" target="_blank" rel="noreferrer">
      <img src={goLogo} class="logo golang" alt="Golang Logo" />
    </a>
  </div>
  <h1>Go + Svelte + Vite</h1>

  <div class="card">
    <Counter />
  </div>

  <p class="read-the-docs">
    Click on the Vite, Svelte and Go logos to learn more
  </p>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .logo.golang:hover {
    filter: drop-shadow(0 0 2em #007d9caa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
