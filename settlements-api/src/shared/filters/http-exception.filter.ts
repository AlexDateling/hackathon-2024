import { Catch, ArgumentsHost, HttpStatus } from '@nestjs/common';
import { BaseExceptionFilter } from '@nestjs/core';

@Catch()
export class HttpExceptionFilter extends BaseExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse();

    let status = HttpStatus.INTERNAL_SERVER_ERROR;

    const message =
      exception instanceof Error ? exception.message : exception.message.error;

    if (exception.status === HttpStatus.NOT_FOUND) {
      status = HttpStatus.NOT_FOUND;
    }

    if (exception.status === HttpStatus.SERVICE_UNAVAILABLE) {
      status = HttpStatus.SERVICE_UNAVAILABLE;
    }

    if (exception.status === HttpStatus.NOT_ACCEPTABLE) {
      status = HttpStatus.NOT_ACCEPTABLE;
    }

    if (exception.status === HttpStatus.EXPECTATION_FAILED) {
      status = HttpStatus.EXPECTATION_FAILED;
    }

    if (exception.status === HttpStatus.BAD_REQUEST) {
      status = HttpStatus.BAD_REQUEST;
    }

    try {
      response.status(status).json({
        status,
        success: false,
        data: [],
        error: message,
        message:
          status === HttpStatus.INTERNAL_SERVER_ERROR
            ? 'Please contact your application administator for any enquiries.'
            : '',
      });
    } catch (error) {
      response.status(status).json({
        status,
        success: false,
        data: [],
        error: error,
        message:
          status === HttpStatus.INTERNAL_SERVER_ERROR
            ? 'Please contact your application administator for any enquiries.'
            : '',
      });
    }
  }
}
