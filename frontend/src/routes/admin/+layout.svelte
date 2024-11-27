<script>
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let isAdmin = false;

  onMount(async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      goto('/auth/login');
      return;
    }

    try {
      const response = await fetch('/api/v1/users/me', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      const data = await response.json();
      
      if (!data.role || data.role !== 'admin') {
        goto('/');
      } else {
        isAdmin = true;
      }
    } catch (error) {
      goto('/auth/login');
    }
  });
</script>

{#if isAdmin}
  <div class="drawer lg:drawer-open">
    <input id="admin-drawer" type="checkbox" class="drawer-toggle" />
    
    <div class="drawer-content flex flex-col">
      <div class="navbar bg-base-100 lg:hidden">
        <div class="flex-none">
          <label for="admin-drawer" class="btn btn-square btn-ghost drawer-button">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-5 h-5 stroke-current">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </label>
        </div>
        <div class="flex-1">
          <span class="text-xl font-bold">Админ панель</span>
        </div>
      </div>
      
      <div class="p-4">
        <slot />
      </div>
    </div>
    
    <div class="drawer-side">
      <label for="admin-drawer" class="drawer-overlay"></label>
      <ul class="menu p-4 w-80 h-full bg-base-200 text-base-content">
        <li class="menu-title">Админ панель</li>
        <li><a href="/admin/products">Товары</a></li>
        <li><a href="/admin/orders">Заказы</a></li>
        <li><a href="/admin/users">Пользователи</a></li>
        <li class="mt-auto"><a href="/" class="text-error">Выйти из админки</a></li>
      </ul>
    </div>
  </div>
{/if} 