<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import Cookies from 'js-cookie';
  import axios from 'axios';

  let user: any = null;
  const token = Cookies.get('token');
  const API = 'http://localhost:3000/api';

  // Column için
  let columnTitle = '';
  let columnDescription = '';
  let columnColor = '#3b82f6';

  // Todo için modal ve inputlar
  let showTodoModal = false;
  let currentColumnIdForTodo = null;
  let todoTitle = '';
  let todoDescription = '';
  let todoColor = '#60a5fa';

  let columns: any[] = [];
  let todos: any[] = [];

  let showColumnModal = false;

  onMount(async () => {
    if (!token) {
      goto('/login');
      return;
    }
    try {
      const res = await axios.post(`${API}/auth/user`, { token });
      if (res.data.status === 'OK') {
        user = res.data.user;
        await fetchData();
      } else {
        goto('/login');
      }
    } catch (e) {
      console.error(e);
      goto('/login');
    }
  });

  async function fetchData() {
    try {
      const [colRes, todoRes] = await Promise.all([
        axios.post(`${API}/column/list`, { token }),
        axios.post(`${API}/todo/list`, { token })
      ]);
      columns = colRes.data.data || [];
      todos = todoRes.data.data || [];
    } catch (e) {
      console.error("Veri çekme hatası:", e);
      columns = [];
      todos = [];
    }
  }

  async function addColumn() {
    if (!columnTitle.trim()) return;
    try {
      await axios.post(`${API}/column/add`, {
        title: columnTitle,
        description: columnDescription,
        color: columnColor,
        user: token
      });
      columnTitle = '';
      columnDescription = '';
      columnColor = '#3b82f6';
      await fetchData();
    } catch (e) {
      console.error("Column ekleme hatası:", e);
    }
  }

  function openColumnModal() {
    columnTitle = '';
    columnDescription = '';
    columnColor = '#3b82f6';
    showColumnModal = true;
  }
  function closeColumnModal() {
    showColumnModal = false;
  }
  async function addColumnModal() {
    if (!columnTitle.trim()) return;
    await addColumn();
    closeColumnModal();
  }

  function openTodoModal(columnId) {
    currentColumnIdForTodo = columnId;
    todoTitle = '';
    todoDescription = '';
    todoColor = '#60a5fa';
    showTodoModal = true;
  }

  function closeTodoModal() {
    showTodoModal = false;
  }

  async function addTodo() {
    if (!todoTitle.trim()) return;

    try {
      console.log('Todo ekleniyor:', { title: todoTitle, path: currentColumnIdForTodo });
      const res = await axios.post(`${API}/todo/add`, {
        title: todoTitle,
        description: todoDescription,
        color: todoColor,
        user: token,
        path: currentColumnIdForTodo
      });
      console.log('Todo eklendi:', res.data);
      await fetchData();
      closeTodoModal();
    } catch (e) {
      console.error("Todo ekleme hatası:", e);
    }
  }

  async function deleteTodo(id: string) {
    try {
      await axios.post(`${API}/todo/del`, { id });
      await fetchData();
    } catch (e) {
      console.error("Todo silme hatası:", e);
    }
  }

  async function deleteColumn(id: string) {
    try {
      await axios.post(`${API}/column/del`, { id });
      await fetchData();
    } catch (e) {
      console.error("Column silme hatası:", e);
    }
  }
</script>

<main class="p-6 max-w-6xl mx-auto space-y-10">

  <!-- Sütunlar ve Todolar -->
  <section>
    <h2 class="text-2xl font-bold mb-4 flex justify-between items-center">
      📦 Sütunlar
      <button
        on:click={openColumnModal}
        class="bg-blue-600 text-white px-3 py-1 rounded text-sm"
      >
        +
      </button>
    </h2>

    <div class="grid md:grid-cols-3 gap-6">
      {#each columns.filter(col => col.user === token) as col}
        <div class="p-4 rounded shadow" style="background-color: {col.color}">
          <div class="flex justify-between items-center mb-2">
            <h3 class="font-bold">{col.title}</h3>
            <button on:click={() => deleteColumn(col.token)} class="text-red-700 font-bold text-xl">×</button>
          </div>
          <p class="text-sm mb-2">{col.description}</p>

          <!-- Todo Listesi -->
          <div class="space-y-2 mt-4">
            {#each todos.filter(todo => todo.path === col.token) as todo}
              <div class="bg-white text-black p-2 rounded shadow flex justify-between items-center"  style="background-color: {todo.color}">
                <div>
                  <div class="font-semibold">{todo.title}</div>
                  <div class="text-xs">{todo.description}</div>
                </div>
                <button on:click={() => deleteTodo(todo.token)} class="text-red-600 font-bold text-xl ml-4">×</button>
              </div>
            {/each}
          </div>

          <!-- Todo Ekle Butonu -->
          <button
            on:click={() => openTodoModal(col.token)}
            class="mt-4 px-3 py-1 bg-green-600 text-white rounded"
          >
            +
          </button>
        </div>
      {/each}
    </div>
  </section>
</main>

<!-- Column Ekleme Modal -->
{#if showColumnModal}
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded p-6 w-96">
      <h3 class="text-xl font-bold mb-4">Yeni Sütun Ekle</h3>

      <input
        type="text"
        placeholder="Başlık"
        bind:value={columnTitle}
        class="w-full p-2 border rounded mb-3"
      />
      <input
        type="text"
        placeholder="Açıklama"
        bind:value={columnDescription}
        class="w-full p-2 border rounded mb-3"
      />
      <label class="block font-semibold mb-1">Renk:</label>
      <input
        type="color"
        bind:value={columnColor}
        class="w-16 h-10 border rounded mb-3"
      />

      <div class="flex justify-end space-x-3">
        <button
          on:click={closeColumnModal}
          class="px-4 py-2 bg-gray-400 rounded text-black"
        >
          İptal
        </button>
        <button
          on:click={addColumnModal}
          class="px-4 py-2 bg-blue-600 text-white rounded"
          disabled={!columnTitle.trim()}
        >
          Ekle
        </button>
      </div>
    </div>
  </div>
{/if}

<!-- Todo Ekleme Modalı -->
{#if showTodoModal}
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded p-6 w-96">
      <h3 class="text-xl font-bold mb-4">Yeni Todo Ekle</h3>

      <input
        type="text"
        placeholder="Başlık"
        bind:value={todoTitle}
        class="w-full p-2 border rounded mb-3"
      />
      <input
        type="text"
        placeholder="Açıklama"
        bind:value={todoDescription}
        class="w-full p-2 border rounded mb-3"
      />
      <label class="block font-semibold mb-1">Renk:</label>
      <input
        type="color"
        bind:value={todoColor}
        class="w-16 h-10 border rounded mb-3"
      />

      <div class="flex justify-end space-x-3">
        <button
          on:click={closeTodoModal}
          class="px-4 py-2 bg-gray-400 rounded text-black"
        >
          İptal
        </button>
        <button
          on:click={addTodo}
          class="px-4 py-2 bg-blue-600 text-white rounded"
          disabled={!todoTitle.trim()}
        >
          Ekle
        </button>
      </div>
    </div>
  </div>
{/if}

<style>
  main {
    font-family: system-ui, sans-serif;
  }
</style>
