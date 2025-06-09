<script lang="ts">
  import { goto } from '$app/navigation';
  import Cookies from 'js-cookie';
  import axios from 'axios';

  let email = '';
  let password = '';
  let username = '';
  let error: string = '';

  const register = async () => {
    const form = { email, password, username };

    try {
      const response = await axios.post('http://localhost:3000/api/auth/register', form);

      if (response.data.status === 'OK') {
        const token = response.data.token;
        Cookies.set('token', token);
        error = 'Kayıt başarılı! Ana sayfaya yönlendiriliyorsunuz...';
        setTimeout(() => goto('/'), 3000);
      } else {
        error = 'Bir hata oluştu. Lütfen bilgilerinizi kontrol edin.';
      }
    } catch (err) {
      console.error('İstek hatası:', err);
      error = 'Sunucuya bağlanılamadı.';
    }
  };
</script>

<main>
  <div class="flex justify-center  items-center">
    <div
      class="mb-1.5 w-full max-w-sm p-4  border-gray-200 rounded-lg shadow sm:p-6 md:p-8 "
    >
      <form class="space-y-6" on:submit|preventDefault={register}>
        <h5 class="text-xl font-medium text-gray-900  ">Register</h5>

        {#if error}
          <div class="bg-green-100 text-green-700 p-4 rounded" role="alert">
            <p class="font-bold">{error}</p>
          </div>
        {/if}

        <div>
          <label for="username" class="block mb-2 text-sm font-medium text-gray-900  ">
            Username
          </label>
          <input
            type="text"
            bind:value={username}
            id="username"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400  "
            placeholder="company"
            required
          />
        </div>

        <div>
          <label for="email" class="block mb-2 text-sm font-medium text-gray-900  ">
            Email
          </label>
          <input
            type="email"
            bind:value={email}
            id="email"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400  "
            placeholder="name@company.com"
            required
          />
        </div>

        <div>
          <label for="password" class="block mb-2 text-sm font-medium text-gray-900  ">
            Password
          </label>
          <input
            type="password"
            bind:value={password}
            id="password"
            placeholder="••••••••"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400  "
            required
          />
        </div>

        <button
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Register
        </button>
      </form>
    </div>
  </div>
</main>
