"use client";
import React, { useState } from "react";

const PostTitle = ({ post }: { post: Blog }) => {
  const [show, setShow] = useState(false);
  return (
    <>
      <li key={post.id}>
        {post.title}{" "}
        <button onClick={() => setShow(!show)}>{show ? "Hide" : "Show"}</button>
      </li>
      {show && <article>Post body: {post.body}</article>}
    </>
  );
};

export default PostTitle;
