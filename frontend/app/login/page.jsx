'use client';

export default function LoginPage() {
  const handleLogin = () => {
    window.location.href = 'http://localhost:8081/oauth/auth/google/login';
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-600 to-indigo-700 px-4">
      <div className="bg-white rounded-2xl shadow-xl p-10 max-w-md w-full text-center">
        <h1 className="text-3xl font-bold text-gray-800 mb-4">Welcome Back!</h1>
        <p className="text-gray-500 mb-8">Login to your account to manage your tasks</p>
        <button
          onClick={handleLogin}
          className="flex items-center justify-center gap-3 bg-red-500 hover:bg-red-600 text-white font-medium py-2 px-4 rounded-lg transition duration-200 w-full"
        >
          <svg
            className="w-5 h-5"
            viewBox="0 0 533.5 544.3"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M533.5 278.4c0-17.4-1.6-34.1-4.7-50.2H272v95h146.8c-6.4 34.6-25.4 63.9-54 83.4v69.3h87.4c51-47 81.3-116.3 81.3-197.5z"
              fill="#4285F4"
            />
            <path
              d="M272 544.3c73.5 0 135.2-24.4 180.3-66.3l-87.4-69.3c-24.2 16.2-55 25.7-92.9 25.7-71 0-131.2-47.9-152.7-112.3H30.7v70.7c45.3 89.3 138.4 151.5 241.3 151.5z"
              fill="#34A853"
            />
            <path
              d="M119.3 321.9c-10.6-31.3-10.6-64.9 0-96.2v-70.7H30.7C-15.1 212.7-15.1 331.7 30.7 418l88.6-69.3z"
              fill="#FBBC05"
            />
            <path
              d="M272 107.7c39.9-.6 78.1 13.9 107.6 40.6l80.4-80.4C411.3 25.5 343.6-2.5 272 0 169.1 0 76 62.2 30.7 151.5l88.6 70.7c21.5-64.4 81.7-112.3 152.7-112.3z"
              fill="#EA4335"
            />
          </svg>
          Continue with Google
        </button>
      </div>
    </div>
  );
}
