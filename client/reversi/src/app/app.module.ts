import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'; // <-- NgModel lives here

import { AngularFontAwesomeModule } from 'angular-font-awesome';

import { AppComponent } from './app.component';
import { RoomsComponent } from './rooms/rooms.component';
import { RoomService } from './room.service';
import { AppRoutingModule } from './/app-routing.module';
import { GameComponent } from './game/game.component';
import { ReversiService } from './reversi.service';
import { WebSocketService } from './websocket.service';


@NgModule({
  declarations: [
    AppComponent,
    RoomsComponent,
    GameComponent
  ],
  imports: [
    AngularFontAwesomeModule,
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [
    RoomService,
    ReversiService,
    WebSocketService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
