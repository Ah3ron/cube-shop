<script>
    import { onMount, onDestroy } from 'svelte';
    import Chart from 'chart.js/auto';
  
    let activeTab = 'catalog';
    let chartCanvas;
    let chart;
    
    // Sample data - Replace with actual API calls
    let products = [];
    let orders = [];
    let isLoading = false;
    let error = null;
    const validStatuses = ['Pending', 'Processing', 'Shipped', 'Delivered'];
  
    // Chart configuration
    const salesData = {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
      datasets: [{
        label: 'Monthly Sales',
        data: [25600, 29400, 31200, 28900, 35600, 42100],
        backgroundColor: 'rgba(59, 130, 246, 0.2)',
        borderColor: 'rgb(59, 130, 246)',
        borderWidth: 2,
        tension: 0.4
      }]
    };
  
    onMount(async () => {
      isLoading = true;
      error = null;
      try {
        // Fetch initial data
        await Promise.all([
          fetchProducts(),
          fetchOrders()
        ]);
      } catch (err) {
        error = err.message;
        console.error('Failed to fetch data:', err);
      } finally {
        isLoading = false;
      }
    });
    async function fetchProducts() {
      // Replace with actual API call
      products = [
        { id: 1, name: 'Product 1', price: 99.99, stock: 50 },
        { id: 2, name: 'Product 2', price: 149.99, stock: 30 }
      ];
    }
  
    async function fetchOrders() {
      // Replace with actual API call
      orders = [
        { id: 1, customer: 'John Doe', date: '2024-01-15', status: 'Pending', total: 299.97 },
        { id: 2, customer: 'Jane Smith', date: '2024-01-14', status: 'Shipped', total: 149.99 }
      ];
    }
  
    async function updateStatus(orderId, newStatus) {
      try {
        // Replace with actual API call
        console.log(`Updating order ${orderId} to ${newStatus}`);
        // await api.updateOrderStatus(orderId, newStatus);
      } catch (error) {
        console.error('Failed to update order status:', error);
      }
    }
  
    function initChart() {
  // Destroy existing chart if it exists
  if (chart) {
    chart.destroy();
    chart = null;
  }

  // Create new chart if we're on stats tab and have the canvas
  if (activeTab === 'stats' && chartCanvas) {
    chart = new Chart(chartCanvas, {
      type: 'line',
      data: salesData,
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: 'top',
          }
        },
        scales: {
          y: {
            beginAtZero: true,
            ticks: {
              callback: (value) => `$${value.toLocaleString()}`
            }
          }
        }
      }
    });
  }
}

// Watch for tab changes and canvas updates
$: if (chartCanvas && activeTab === 'stats') {
  // Use setTimeout to ensure DOM is ready
  setTimeout(() => {
    initChart();
  }, 0);
}

onDestroy(() => {
  if (chart) {
    chart.destroy();
    chart = null;
  }
});
  </script>
  
  <div class="container mx-auto px-3 py-8">
    <h1 class="text-3xl font-bold mb-8">Admin Dashboard</h1>
  
    <div class="tabs tabs-boxed mb-6">
      <button 
        class="tab {activeTab === 'catalog' ? 'tab-active' : ''}"
        on:click={() => activeTab = 'catalog'}>
        Catalog Management
      </button>
      <button 
        class="tab {activeTab === 'orders' ? 'tab-active' : ''}"
        on:click={() => activeTab = 'orders'}>
        Orders
      </button>
      <button 
        class="tab {activeTab === 'stats' ? 'tab-active' : ''}"
        on:click={() => activeTab = 'stats'}>
        Statistics
      </button>
    </div>

    {#if isLoading}
      <div class="loading">Loading data...</div>
    {:else if error}
      <div class="error">Error: {error}</div>
    {:else}
      {#if activeTab === 'catalog'}
        <div class="card bg-base-100 shadow">
          <div class="card-body">
            <h2 class="card-title">Product Management</h2>
            <div class="overflow-x-auto">
              <table class="table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Price</th>
                    <th>Stock</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {#each products as product}
                    <tr>
                      <td>{product.id}</td>
                      <td>{product.name}</td>
                      <td>${product.price}</td>
                      <td>{product.stock}</td>
                      <td>
                        <button class="btn btn-sm btn-outline">Edit</button>
                        <button class="btn btn-sm btn-error ml-2">Delete</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
            <button class="btn btn-primary mt-4">Add New Product</button>
          </div>
        </div>
      {/if}
    
      {#if activeTab === 'orders'}
        <div class="card bg-base-100 shadow">
          <div class="card-body">
            <h2 class="card-title">Order Management</h2>
            <div class="overflow-x-auto">
              <table class="table">
                <thead>
                  <tr>
                    <th>Order ID</th>
                    <th>Customer</th>
                    <th>Date</th>
                    <th>Status</th>
                    <th>Total</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {#each orders as order}
                    <tr>
                      <td>{order.id}</td>
                      <td>{order.customer}</td>
                      <td>{order.date}</td>
                      <td>
                        <select 
                          class="select select-bordered select-sm"
                          on:change={(e) => updateStatus(order.id, e.target.value)}
                        >
                          {#each validStatuses as status}
                            <option>{status}</option>
                          {/each}
                        </select>
                      </td>
                      <td>${order.total}</td>
                      <td>
                        <button class="btn btn-sm btn-outline">View</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      {/if}
    
      {#if activeTab === 'stats'}
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">Sales Overview</h2>
              <canvas bind:this={chartCanvas}></canvas>
            </div>
          </div>
          
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">Quick Stats</h2>
              <div class="stats stats-vertical shadow">
                <div class="stat">
                  <div class="stat-title">Total Sales</div>
                  <div class="stat-value">$25,600</div>
                  <div class="stat-desc">21% more than last month</div>
                </div>
                
                <div class="stat">
                  <div class="stat-title">New Orders</div>
                  <div class="stat-value">578</div>
                  <div class="stat-desc">↗︎ 14 (2%)</div>
                </div>
                
                <div class="stat">
                  <div class="stat-title">Active Products</div>
                  <div class="stat-value">126</div>
                  <div class="stat-desc">↘︎ 90 (14%)</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      {/if}
    {/if}
  </div>
  
  <style>
    :global(.chart-container) {
      position: relative;
      height: 300px;
      width: 100%;
    }
  </style>