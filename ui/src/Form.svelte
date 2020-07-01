<script>
  import store from "./store";

  let offset = 0;
  let limit = 10;

  function loadProducts() {
    store.loadProducts({ offset, limit });
  }

  $: isOffsetValid =
    offset === undefined || (Number.isInteger(offset) && 0 <= offset);

  $: isLimitValid =
    limit === undefined ||
    (Number.isInteger(limit) && 1 <= limit && limit <= 100);

  $: isFormValid = isOffsetValid && isLimitValid;

  $: canLoadProducts = !$store.isLoading && isFormValid;
</script>

<form>
  <div class="form-row">
    <div class="form-group col-md-2">
      <label for="inputOffset">Offset</label>
      <input
        type="number"
        min="0"
        class="form-control"
        class:is-invalid={!isOffsetValid}
        id="inputOffset"
        bind:value={offset} />
      <small
        class={isOffsetValid ? 'form-text text-muted' : 'invalid-feedback'}>
        Type integer greater than or equal to zero.
      </small>
    </div>
    <div class="form-group col-md-2">
      <label for="inputLimit">Limit</label>
      <input
        type="number"
        min="1"
        max="100"
        class="form-control"
        class:is-invalid={!isLimitValid}
        id="inputLimit"
        bind:value={limit} />
      <small class={isLimitValid ? 'form-text text-muted' : 'invalid-feedback'}>
        Type integer between 1 and 100.
      </small>
    </div>
  </div>
  <button
    type="button"
    class="btn btn-primary"
    on:click={loadProducts}
    disabled={!canLoadProducts}>
    Load products
  </button>
</form>
