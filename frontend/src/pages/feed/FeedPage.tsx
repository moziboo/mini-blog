// filepath: /Users/stephen.garrett/Dev/mini-blog/frontend/src/pages/feed/FeedPage.tsx
import { useState } from 'react';
import { PostCard } from '../../components/blog/PostCard';
import { usePosts } from '../../hooks/usePosts';

export function FeedPage() {
  const { posts, loading, error } = usePosts();
  const [searchTerm, setSearchTerm] = useState('');

  const filteredPosts = posts.filter(
    (post) =>
      post.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      post.content.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (post.author &&
        post.author.toLowerCase().includes(searchTerm.toLowerCase()))
  );

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <h1 className="text-3xl font-bold text-gray-900">Feed</h1>

        <div className="w-full sm:w-auto">
          <input
            type="text"
            placeholder="Search posts..."
            className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
      </div>

      {loading ? (
        <div className="flex justify-center items-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
        </div>
      ) : error ? (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
          <p>Error loading posts. Please try again later.</p>
        </div>
      ) : filteredPosts.length === 0 ? (
        <div className="bg-gray-50 border border-gray-200 text-gray-700 px-4 py-5 rounded text-center">
          {searchTerm ? (
            <p>
              No posts match your search criteria. Try a different search term.
            </p>
          ) : (
            <p>No posts found. Check back later for new content!</p>
          )}
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-8">
          {filteredPosts.map((post) => (
            <PostCard key={post.id} post={post} />
          ))}
        </div>
      )}

      {filteredPosts.length > 0 && (
        <div className="pt-4 text-gray-500 text-sm">
          Showing {filteredPosts.length}{' '}
          {filteredPosts.length === 1 ? 'post' : 'posts'}
          {searchTerm && ` matching "${searchTerm}"`}
        </div>
      )}
    </div>
  );
}
