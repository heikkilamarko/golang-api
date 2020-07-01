<script>
  export let product;
</script>

<style>
  .card {
    margin-bottom: 1rem;
  }

  .card:hover {
    border-color: rgba(0, 123, 255, 0.5);
    background-color: rgba(0, 123, 255, 0.1);
  }

  dl {
    margin-bottom: 0;
  }

  dd:last-child {
    margin-bottom: 0;
  }

  @media (min-width: 768px) {
    dt {
      text-align: right;
    }
  }
</style>

{#if product}
  <div class="card p-3">
    <dl class="row">
      <dt class="col-md-2">Name</dt>
      <dd class="col-md-10">{product.name}</dd>
      <dt class="col-md-2">Description</dt>
      <dd class="col-md-10">{product.description || '-'}</dd>
      <dt class="col-md-2">Price</dt>
      <dd class="col-md-10">$ {product.price}</dd>
      <dt class="col-md-2">Comment</dt>
      <dd class="col-md-10">{product.comment || '-'}</dd>
    </dl>
  </div>
{/if}
