import { Injectable, NotFoundException } from '@nestjs/common';
import { Todo, TodoStatus } from './models/todo.models';

@Injectable()
export class TodoService {
  // 今回はDBと接続しないのでメモリ上にTodoを保存します。
  private todos: Todo[] = [];

  // 全件取得のメソッド
  findAll(): Todo[] {
    const todo = new Todo();
    todo.id = '1';
    todo.title = 'title';
    todo.description = 'desc';
    todo.status = TodoStatus.NEW;
    todo.createdAt = new Date('2021-01-01 00:00:00');
    todo.updatedAt = new Date('2021-01-02 00:00:00');

    this.todos[0] = todo;

    return this.todos;
  }

  // idを元に一件取得のメソッド
  findOneById(id: string): Todo {
    const todo = new Todo();
    todo.id = '1';
    todo.title = 'title';
    todo.description = 'desc';
    todo.status = TodoStatus.NEW;
    todo.createdAt = new Date('2021-01-01 00:00:00');
    todo.updatedAt = new Date('2021-01-02 00:00:00');

    this.todos[0] = todo;

    const result = this.todos.find((todo) => id === todo.id);
    if (!result) {
      // なかったら404エラーを返す。ビルトインのエラーも豊富にあってエラー処理も結構楽
      // https://docs.nestjs.com/exception-filters#built-in-http-exceptions
      throw new NotFoundException();
    }
    return result;
  }
}
