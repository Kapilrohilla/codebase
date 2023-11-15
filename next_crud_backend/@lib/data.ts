type POST = {
  id: string;
  title: string;
  desc: string;
  date: Date;
};
let posts: POST[] = [];

// handlers
export const getPosts = () => posts;
export const addPost = (post: POST) => {
  posts.push(post);
};
export const deletePost = (id: string) => {
  posts = posts.filter((post) => {
    return post.id !== id;
  });
};
export const updatePost = (id: string, title: string, desc: string) => {
  const post = posts.find((post) => (post.id = id));

  if (post) {
    post.title = title;
    post.desc = desc;
  } else {
    throw new Error(`NO POST FOUND by ID = ${id}`);
  }
};

export const getById = (id: string) => {
  const post = posts.find((post) => post.id === id);
  if (!post) {
    throw new Error(`No Post Found with ID = ${id}`);
  } else {
    return post;
  }
};
