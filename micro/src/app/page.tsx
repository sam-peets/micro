"use client";

import { useEffect, useState } from 'react';
import { GetRecent } from './lib/api';
import { Post } from './types';
import PostList from './components/postList';
import CreatePost from './components/createPost';

// page content
export default function Page() {
  const [data, setData] = useState<Post[] | []>([]);

  useEffect(() => {
    GetRecent(10, 0).then(x => {
      setData(x)
    })
  }, [])

  return (
    <main>
      <h1>Micro</h1>
      <CreatePost action={(p: Post) => setData([p, ...data])}/>
      <hr style={{border: "none", height: "1px", color: "lightgray", backgroundColor: "lightgray"}}/>
      <PostList posts={data}></PostList>
    </main>
  );
}