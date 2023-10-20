import getBlogs from "@/lib/getUserBlogs";
import Link from "next/link";
import React from "react";

const UserBlogs = async ({ params }: { params: { userId: string } }) => {
  const blogs = await getBlogs(params.userId);
  return (
    <>
      <h1>Blogs: {blogs.length} </h1>
      <ul>
        {blogs.map((blog) => {
          return (
            <li key={blog.id}>
              {blog.title}
              <Link href={`/users/${params.userId}/posts`}>
                <button>Goto this blog</button>
              </Link>
            </li>
          );
        })}
      </ul>
    </>
  );
};

export default UserBlogs;
