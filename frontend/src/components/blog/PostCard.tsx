import { Link } from 'react-router-dom';
import { Card } from '../ui/Card';

export interface Post {
  id: string;
  title: string;
  content: string;
  excerpt?: string;
  author?: string;
  publishedAt?: string;
  slug?: string;
}

interface PostCardProps {
  post: Post;
}

export function PostCard({ post }: PostCardProps) {
  const {
    id,
    title,
    content,
    excerpt,
    publishedAt,
    slug,
  } = post;

  const formattedDate = publishedAt
    ? new Date(publishedAt).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      })
    : null;

  const postLink = slug ? `/post/${slug}` : `/post/${id}`;
  const summary =
    excerpt || content.substring(0, 150) + (content.length > 150 ? '...' : '');

  return (
    <Card
      variant="bordered"
      className="hover:shadow-md transition-shadow duration-200"
    >
      <div className="space-y-3">
        <Link to={postLink}>
          <h3 className="text-xl font-semibold text-gray-900 hover:text-blue-600">
            {title}
          </h3>
        </Link>

        <div className="flex items-center text-sm text-gray-500">
          {formattedDate && (
            <>
              <span className="mx-1">•</span>
              <time dateTime={publishedAt}>{formattedDate}</time>
            </>
          )}
        </div>

        <p className="text-gray-600">{summary}</p>

        <div className="pt-2">
          <Link
            to={postLink}
            className="text-blue-600 hover:text-blue-800 font-medium text-sm"
          >
            Read more →
          </Link>
        </div>
      </div>
    </Card>
  );
}
