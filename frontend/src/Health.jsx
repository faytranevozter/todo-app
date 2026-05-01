import { useState, useEffect } from 'react';

export default function Health() {
  const [status, setStatus] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const controller = new AbortController();
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';
    
    fetch(`${apiUrl}/api/health`, { signal: controller.signal })
      .then(res => {
        if (!res.ok) {
          throw new Error(`Backend returned ${res.status}`);
        }
        return res.json();
      })
      .then(data => {
        setStatus(data.status);
        setLoading(false);
      })
      .catch(err => {
        if (err.name !== 'AbortError') {
          setError(err.message);
          setLoading(false);
        }
      });

    return () => controller.abort();
  }, []);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="bg-white rounded-lg shadow-lg p-8 max-w-md w-full">
        <h1 className="text-3xl font-bold text-gray-800 mb-6">Health Check</h1>
        
        {loading && (
          <div className="text-center">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
            <p className="mt-4 text-gray-600">Checking backend health...</p>
          </div>
        )}
        
        {status && !loading && (
          <div className="text-center">
            <div className="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded mb-4">
              ✓ Backend is healthy
            </div>
            <p className="text-gray-600">Status: <code className="bg-gray-100 px-2 py-1 rounded text-sm font-mono">{status}</code></p>
          </div>
        )}
        
        {error && !loading && (
          <div className="text-center">
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              ✗ Failed to connect
            </div>
            <p className="text-gray-600 text-sm">{error}</p>
          </div>
        )}
      </div>
    </div>
  );
}
