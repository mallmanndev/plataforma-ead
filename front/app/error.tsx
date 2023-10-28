"use client";

export default function ErrorPage() {
  return (
    <div className="bg-red-500 h-screen flex flex-col justify-center items-center">
      <div className="text-white text-4xl font-semibold mb-4">
        Oops! Algo deu errado.
      </div>
      <p className="text-white text-xl mb-8">
        A página que você está procurando não pôde ser encontrada.
      </p>
      <a
        href="/home"
        className="bg-white text-red-500 hover:bg-red-500 hover:text-white py-2 px-4 rounded-full text-lg font-semibold"
      >
        Voltar à Página Inicial
      </a>
    </div>
  );
}
