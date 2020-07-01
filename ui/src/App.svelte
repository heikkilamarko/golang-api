<script>
  import { onMount } from "svelte";
  import store from "./store";
  import Form from "./Form.svelte";
  import ProductList from "./ProductList.svelte";

  onMount(() => {
    store.loadProducts();
  });

  $: products = $store.products.data || [];
</script>

<Form />
<hr />
<ProductList {products} />
