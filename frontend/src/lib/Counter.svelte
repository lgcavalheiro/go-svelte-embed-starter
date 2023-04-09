<script lang="ts">
  interface ResponseDouble extends Response {
    result?: number;
  }

  let count: number = 0;
  let doubleCount: number = 0;

  const increment = () => {
    count += 1;
  };

  const getDouble = async () => {
    const response: ResponseDouble = await fetch(`/api/double?number=${count}`);
    const jsonRes = await response.json();
    doubleCount = jsonRes.result;
  };
</script>

<button on:click={increment} data-testid="counter-btn">
  count is {count}
</button>

<button on:click={getDouble} data-testid="double-btn">
  Gopher doubles the counter: {doubleCount}
</button>
