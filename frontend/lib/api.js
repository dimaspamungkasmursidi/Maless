export async function fetchFromAPI(path, options = {}) {
    return fetch(`http://localhost:8081${path}`, {
      ...options,
      credentials: 'include',
    });
  }
  