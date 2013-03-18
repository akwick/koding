
class LandingPageSideBar extends KDView

  constructor:(isLoggedIn = no)->

    options     =
      lazyDomId : 'landing-page-sidebar'

    super options

    @navController = new LandingPageNavigationController
      view         : new NavigationList
        itemClass  : LandingPageNavigationLink
        type       : "navigation"
      scrollView   : no
      wrapper      : no
      delegate     : @
    ,
      items : [
        { title : "Register", action : "register", loggedOut : yes }
        { type  : "separator" }
        { title : "Logout",   action : "logout",   loggedIn  : yes }
        { title : "Login",    action : "login",    loggedOut : yes }
      ]

    @on 'ListItemsInstantiated' , =>
      $("#profile-static-nav").remove()

    @addSubView @nav = @navController.getView()

class LandingPageNavigationController extends NavigationController

  constructor: ->
    super

    @lc = @getSingleton 'lazyDomController'

  instantiateListItems:(items)->

    # Build groups menu
    if @lc.userEnteredFromGroup()

      {groupEntryPoint} = KD.config

      if KD.isLoggedIn()
        KD.whoami().fetchGroupRoles groupEntryPoint, (err, roles)=>
          if err then console.warn err
          else if roles.length
            items.unshift \
              { title: 'Open Group', path: "/#{if groupEntryPoint is 'koding' then '' else groupEntryPoint+'/'}Activity"}
            @_instantiateListItems items
          else
            KD.remote.api.JMembershipPolicy.byGroupSlug groupEntryPoint,
              (err, policy)=>
                if err then console.warn err
                else if policy?.approvalEnabled
                  items.unshift \
                    { title: 'Request access', action: 'request'}
                else
                  items.unshift \
                    { title: 'Join Group', action: 'join-group'}
                @_instantiateListItems items

      else
        items.unshift { title: 'Request access', action: 'request'}

        if groupEntryPoint is "koding" then items.first.title = "Request Invite"

        @_instantiateListItems items

    else if @lc.userEnteredFromProfile()

      log 'entered from profile!'
      profileItems = [
        { title : 'My Activities', action : 'activity', type : 'user'}
        { title : 'My Topics',     action : 'topics',   type : 'user' }
        { title : 'My People',     action : 'members',  type : 'user'}
        { title : 'My Groups',     action : 'groups',   type : 'user'}
        { title : 'My Apps',       action : 'apps',     type : 'user'}
        { type  : "separator" }
      ]
      items = profileItems.concat items
      @_instantiateListItems items
    else
      @_instantiateListItems items

  _instantiateListItems:(items)->
    @getDelegate().emit 'ListItemsInstantiated'
    newItems = for itemData in items
      if KD.isLoggedIn()
        continue if itemData.loggedOut
      else
        continue if itemData.loggedIn
      @getListView().addItem itemData

fetchGroupFirst = (callback)->
  {groupEntryPoint} = KD.config
  KD.remote.api.JGroup.one slug: groupEntryPoint, (err, group)=>
    error err if err
    if err then new KDNotificationView
      title : "An error occured, please try again"
    else unless group?
      new KDNotificationView title : "No such group!"
    else callback group

class LandingPageNavigationLink extends NavigationLink

  constructor:(options = {}, data)->
    data.type or= "account"
    super options, data
    @lc = @getSingleton 'lazyDomController'

  openPath:(path)->
    @getSingleton('router').handleRoute path
    @getSingleton('lazyDomController').hideLandingPage()

  click:(event)->
    {action, appPath, title, path, type} = @getData()
    log "here", @getData()

    mc = @getSingleton 'mainController'
    {loginScreen} = mc

    if path
      @openPath path
      return

    {groupEntryPoint, profileEntryPoint} = KD.config

    switch action
      when 'login'
        loginScreen.animateToForm 'login'
      when 'register'
        loginScreen.animateToForm 'register'
      when 'request'
        if KD.isLoggedIn()
          fetchGroupFirst (group)=>
            @getSingleton('groupsController').openPrivateGroup group
        else
          loginScreen.animateToForm 'lr'
      when 'join-group'
        fetchGroupFirst (group)=>
          group.join (err, response)=>
            error err if err
            if err then new KDNotificationView
              title : "An error occured, please try again"
            else
              new KDNotificationView
                title : "You successfully joined to group!"
              @openPath "/#{groupEntryPoint}/Activity"

      when 'logout'
        mainController = @getSingleton('mainController')
        mainController.mainViewController.getView().hide()
        @openPath '/Logout'

      when 'activity'
        @openPath "/#{profileEntryPoint}/Activity"
        @lc.hideLandingPage()
        log 'Activity'
      when 'topics'
        @openPath "/#{profileEntryPoint}/Topics"
        @lc.hideLandingPage()

        log 'Topics'
      when 'members'
        @openPath "/#{profileEntryPoint}/Members"
        @lc.hideLandingPage()
        log 'Members'
      when 'groups'
        @openPath "/#{profileEntryPoint}/Groups"
        @lc.hideLandingPage()
        log 'Groups'
      when 'apps'
        @openPath "/#{profileEntryPoint}/Apps"
        @lc.hideLandingPage()
        log 'Apps'

