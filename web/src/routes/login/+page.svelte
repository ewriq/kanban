<script lang="ts">
  import { goto } from '$app/navigation';
  import Cookies from 'js-cookie';
  import axios from 'axios';

  let email = '';
  let password = '';
  let error: string = '';

  const login = async () => {
    const form = { email, password };

    try {
      const response = await axios.post('http://localhost:3000/api/auth/login', form);
      if (response.data.status === 'OK') {
        const token = response.data.token;
        Cookies.set('token', token);
        error = 'Başarılı! Yönlendiriliyorsunuz...';
        setTimeout(() => goto('/dashboard'), 3000); // SvelteKit yönlendirme
      } else {
        error = 'Email ya da parola hatalı';
      }
    } catch (err) {
      console.error('Hata oluştu:', err);
      error = 'Sunucu hatası oluştu.';
    }
  };
</script>

<main>
  <div class="flex justify-center  items-center">
    <div
      class="mb-1.5 w-full max-w-sm p-4  border-gray-200 rounded-lg shadow sm:p-6 md:p-8 "
    >
      <form class="space-y-6" on:submit|preventDefault={login}>
        <h5 class="text-xl font-medium text-gray-900 ">Login</h5>

        {#if error}
          <div class="bg-green-100 text-green-700 p-4 rounded" role="alert">
            <p class="font-bold">{error}</p>
          </div>
        {/if}

        <div>
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900 ">
            Your email
          </label>
          <input
            type="email"
            bind:value={email}
            id="email"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 "
            placeholder="name@company.com"
            required
          />
        </div>

        <div>
          <label
            for="password"
            class="block mb-2 text-sm font-medium text-gray-900 "
          >
            Your password
          </label>
          <input
            type="password"
            bind:value={password}
            id="password"
            placeholder="••••••••"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 "
            required
          />
        </div>

        <button
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Login
        </button>
      </form>
    </div>
  </div>
</main>
