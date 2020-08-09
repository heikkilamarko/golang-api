<script>
  import store from "./store";
  import ProductCard from "./ProductCard.svelte";

  $: products = $store.products.data || [];
</script>

{#each products as product (product.id)}
  <ProductCard {product} />
{/each}
