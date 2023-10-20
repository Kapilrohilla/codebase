export default async function getUserBlogs(userId: string): Promise<Blog[]> {
  const res = await fetch(
    `https://jsonplaceholder.typicode.com/posts?userId=${userId}`
  );
  if (!res.ok) throw new Error("Failed to fetch UserData");
  return res.json();
}
