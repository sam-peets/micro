'use client'
 
import { useSearchParams } from 'next/navigation'
import { GetUser, GetUserPosts } from '../lib/api'
import { Post, User } from '../types'
import { useEffect, useState } from 'react'
import PostList from '../components/postList'
 
export default function Page() {
  const [user, setUser] = useState<User | null>(null)
  const [posts, setPosts] = useState<Post[] | []>([])

  const searchParams = useSearchParams()
  const uid = searchParams.get('id')
  if (!uid) {
    return;
  }

  useEffect(() => {
    GetUser(parseInt(uid)).then(u => {
      setUser(u)
    })
  }, [])

  useEffect(() => {
    GetUserPosts(parseInt(uid), 10, 0).then(p => {
      setPosts(p)
    })
  }, [])
  
  return (
    <main>
      <h1>{user?.username}</h1>
      <PostList posts={posts}/>

    </main>
  )
}