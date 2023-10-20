export default async function getAllBlogs(): Promise<Blog[]> {
  const res = await fetch("https://jsonplaceholder.typicode.com/posts");
  if (!res.ok) throw new Error("failed to fetch data");
  return res.json();
}
