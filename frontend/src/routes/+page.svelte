// src/routes/+page.svelte
<script>
  import { onMount } from 'svelte';

  let products = [];

  onMount(async () => {
    try {
      const response = await fetch('/api/products');
      products = await response.json();
    } catch (error) {
      console.error('Failed to fetch products:', error);
    }
  });
</script>

<main>
  <h1>Welcome to CubeShop</h1>
  
  <section class="products">
    {#if products.length > 0}
      {#each products as product}
        <div class="product-card">
          <img src={product.image} alt={product.name} />
          <h2>{product.name}</h2>
          <p>{product.price}</p>
          <button>Add to Cart</button>
        </div>
      {/each}
    {:else}
      <p>Loading products...</p>
    {/if}
  </section>
</main>

<style>
  main {
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .products {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 2rem;
    margin-top: 2rem;
  }

  .product-card {
    border: 1px solid #ddd;
    padding: 1rem;
    border-radius: 8px;
    text-align: center;
  }

  .product-card img {
    max-width: 100%;
    height: auto;
  }

  button {
    background: #4CAF50;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    cursor: pointer;
  }

  button:hover {
    background: #45a049;
  }
</style>