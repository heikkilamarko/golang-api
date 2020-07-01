<script>
  import ProductCard from "./ProductCard.svelte";

  export let products = [];
</script>

{#each products as product (product.id)}
  <ProductCard {product} />
{/each}
