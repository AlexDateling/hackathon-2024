import { Logger } from '@nestjs/common';

export class ApplicationLogger extends Logger {
    error(message: string, trace: string) {
      super.error(message, trace);
    }

    log(message: string) {
      super.error(message);
    }

    warn(message: string) {
      super.error(message);
    }

    debug(message: string) {
      super.error(message);
    }

    verbose(message: string) {
      super.error(message);
    }
}
