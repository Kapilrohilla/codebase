import getAllBlogs from "@/lib/getAllBlogs";
import React from "react";

const Posts = async () => {
  const blogs = await getAllBlogs();

  return (
    <>
      <h1>All Posts</h1>
      <ul>
        {blogs.map((blog) => {
          return <li key={blog.id}>{blog.title}</li>;
        })}
      </ul>
    </>
  );
};

export default Posts;
