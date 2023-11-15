import { deletePost, getById, updatePost } from "@/@lib/data";
import { NextResponse } from "next/server";

export const GET = async (req: Request) => {
  // console.log("GET");
  // get post by id
  const id = req.url.split("blogs/")[1];
  try {
    const post = getById(id);
    return NextResponse.json(post, { status: 200 });
  } catch (err) {
    // console.log(error);
    return NextResponse.json({ err }, { status: 500 });
  }
};

export const PUT = async (req: Request) => {
  const { title, desc } = await req.json();
  const id = req.url.split("blogs/")[1];

  try {
    updatePost(id, title, desc);
    return NextResponse.json("data updated", { status: 200 });
  } catch (err) {
    return NextResponse.json({ err }, { status: 500 });
  }
};

export const DELETE = async (req: Request, res: Response) => {
  const id = req.url.split("blogs/")[1];

  try {
    deletePost(id);
    return NextResponse.json({ msg: "successful" }, { status: 200 });
  } catch (err) {
    return NextResponse.json({ msg: err }, { status: 500 });
  }
};
