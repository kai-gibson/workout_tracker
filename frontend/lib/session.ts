function setSessionStorage<T>(key: string, value: T): void {
  if (typeof window === 'undefined' || !window.sessionStorage) {
    throw new Error(
      'Not running in a browser environment or sessionStorage is not available.'
    );
  }

  window.sessionStorage.setItem(key, JSON.stringify(value));
}

function getSessionStorage<T>(key: string): T | null {
  if (typeof window === 'undefined' || !window.sessionStorage) {
    throw new Error(
      'Not running in a browser environment or sessionStorage is not available.'
    );
  }

  const value = window.sessionStorage.getItem(key);
  try {
    return value ? JSON.parse(value) : null;
  } catch (error) {
    console.warn(`Error parsing sessionStorage value for key "${key}":`, error);
    return null;
  }
}

export {setSessionStorage, getSessionStorage};
