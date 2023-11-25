export default async function Home() {
  const apiUrl = process.env.API_URL;
  if (!apiUrl) {
    throw new Error("API_URL is not defined in the environment variables");
  }

  const data = await fetch(apiUrl);
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 text-5xl">
      Top
    </main>
  );
}
