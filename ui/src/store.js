import { writable } from "svelte/store";
import produce from "immer";
import axios from "axios";
import { alertError } from "./utils";
import { API_URL } from "./config";

function createStore(initialValue) {
  const store = writable(initialValue);

  const { subscribe, update } = store;

  async function loadProducts(query, apiKey) {
    query = query || {};

    setLoading(true);

    try {
      const response = await axios.get(`${API_URL}/products`, {
        headers: { "X-Api-Key": apiKey },
        params: { ...query },
      });
      setProducts(response.data);
    } catch (error) {
      alertError(error);
    }

    setLoading(false);
  }

  function setLoading(isLoading = true) {
    setState((d) => (d.isLoading = isLoading));
  }

  function setProducts(products) {
    setState((d) => (d.products = products));
  }

  function setState(fn) {
    update(
      produce((d) => {
        fn(d);
      })
    );
  }

  return {
    subscribe,
    loadProducts,
  };
}

const initialValue = {
  isLoading: false,
  products: {},
};

export default createStore(initialValue);
